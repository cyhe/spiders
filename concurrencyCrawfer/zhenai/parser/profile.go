package parser

import (
	"spiders/concurrencyCrawfer/engine"
	"regexp"
	"strconv"
	"spiders/concurrencyCrawfer/model"
	"spiders/concurrencyCrawfer/config"
)

//const ageRe = `<td><span class = "label">年龄：</span>([\d]+岁)</td>`
//const marriageRe = `<td><span class="label">婚况：</span>([^<]+)</td>`

var (
	ageRe        = regexp.MustCompile(`<td><span class="label">年龄：</span>(\d+)岁</td>`)
	marriageRe   = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
	nameRe       = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
	genderRe     = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
	heightRe     = regexp.MustCompile(`<td><span class="label">身高：</span>(\d+)CM</td>`)
	weightRe     = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">(\d+)KG</span></td>`)
	incomeRe     = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
	educationRe  = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
	occupationRe = regexp.MustCompile(`<td><span class="label">职业：</span><span field="">([^<]+)</span></td>`)
	hokouRe      = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
	xinzuoRe     = regexp.MustCompile(`<td><span class="label">星座：</span><span field="">([^<]+)</span></td>`)
	houseRe      = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
	carRe        = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)
	idUrlRe      = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)
	guessRe      = regexp.MustCompile(`<a class="exp-user-name"[^>]*href="(http://album.zhenai.com/u/[\d]+)">([^<]+)</a>`)
)

func ParseProfile(contents []byte, name string, url string) engine.ParserResult {

	profile := model.Profile{}

	profile.Age = extractInt(contents, ageRe)
	profile.Marriage = extractString(contents, marriageRe)
	profile.Gender = extractString(contents, genderRe)
	profile.Height = extractInt(contents, heightRe)
	profile.Weight = extractInt(contents, weightRe)
	profile.Income = extractString(contents, incomeRe)
	profile.Education = extractString(contents, educationRe)
	profile.Occupation = extractString(contents, occupationRe)
	profile.Hokou = extractString(contents, hokouRe)
	profile.Xinzuo = extractString(contents, xinzuoRe)
	profile.House = extractString(contents, houseRe)
	profile.Car = extractString(contents, carRe)

	profile.Name = name

	result := engine.ParserResult{
		Items: []engine.Item{
			{
				Url:     url,
				Type:    config.TypeName,
				Id:      extractString([]byte(url), idUrlRe),
				Payload: profile,
			},
		},
	}

	matches := guessRe.FindAllSubmatch(contents, -1)

	for _, match := range matches {
		result.Requests = append(result.Requests,
			engine.Request{
				Url:        string(match[1]),
				ParserFunc: ProfileParser(string(match[2])),
			},
		)
	}

	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}

func extractInt(contents []byte, re *regexp.Regexp) int {
	num, err := strconv.Atoi(extractString(contents, re))
	if err != nil {
		return 0
	}
	return num
}

func ProfileParser(name string) engine.ParserFunc {
	return func(c []byte,url string) engine.ParserResult {
		return ParseProfile(c, name, url)
	}
}

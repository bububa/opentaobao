package wenyuvideo

import (
	"sync"
)

// PersonDetailRbo 结构体
type PersonDetailRbo struct {
	// 人物类型：unknow,director,starring,performer,singer,lyricswriter,composer,screenwriter,producer,host,voice,      * executive_producer,teacher,original, interview,paike
	PersonTypeList []string `json:"person_type_list,omitempty" xml:"person_type_list>string,omitempty"`
	// 人物国籍
	NationalityList []string `json:"nationality_list,omitempty" xml:"nationality_list>string,omitempty"`
	// 职业
	OccupationList []string `json:"occupation_list,omitempty" xml:"occupation_list>string,omitempty"`
	// 人物性质：牛人,拍客,网络红人
	PersonKindList []string `json:"person_kind_list,omitempty" xml:"person_kind_list>string,omitempty"`
	// 组合成员，如果人物是一个组合，则允许设置下面的成员，人员也是一人物记录。输出格式[{&#39;id&#39;:20349,&#39;name&#39;:&#39;张宇凡&#39;},{&#39;id&#39;:35994,&#39;name&#39;:&#39;袁泉&#39;}]
	MemberList []string `json:"member_list,omitempty" xml:"member_list>string,omitempty"`
	// 人物相关视频推荐
	RefShows []ShowBaseRbo `json:"ref_shows,omitempty" xml:"ref_shows>show_base_rbo,omitempty"`
	// 人物相关人物推荐
	RefPersons []PersonRbo `json:"ref_persons,omitempty" xml:"ref_persons>person_rbo,omitempty"`
	// 人物名字
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 图片
	ThumbUrl string `json:"thumb_url,omitempty" xml:"thumb_url,omitempty"`
	// 300*300人物头像
	ThumbUrlLg string `json:"thumb_url_lg,omitempty" xml:"thumb_url_lg,omitempty"`
	// 人物海报
	PosterUrl string `json:"poster_url,omitempty" xml:"poster_url,omitempty"`
	// 人物写真
	PosterUrlH string `json:"poster_url_h,omitempty" xml:"poster_url_h,omitempty"`
	// 简介
	PersonDesc string `json:"person_desc,omitempty" xml:"person_desc,omitempty"`
	// 性别 U(nknow):未知 M(ale):男 F(emale):女 G(roup):组合
	Gender string `json:"gender,omitempty" xml:"gender,omitempty"`
	// 生日
	Birthday string `json:"birthday,omitempty" xml:"birthday,omitempty"`
	// 忌日
	Deathday string `json:"deathday,omitempty" xml:"deathday,omitempty"`
	// 出生地
	Homeplace string `json:"homeplace,omitempty" xml:"homeplace,omitempty"`
	// 身高，单位cm
	Height string `json:"height,omitempty" xml:"height,omitempty"`
	// 血型 A、B、AB、O
	BloodType string `json:"blood_type,omitempty" xml:"blood_type,omitempty"`
	// 星座
	Constellation string `json:"constellation,omitempty" xml:"constellation,omitempty"`
	// 主键(优酷人物ID)
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolPersonDetailRbo = sync.Pool{
	New: func() any {
		return new(PersonDetailRbo)
	},
}

// GetPersonDetailRbo() 从对象池中获取PersonDetailRbo
func GetPersonDetailRbo() *PersonDetailRbo {
	return poolPersonDetailRbo.Get().(*PersonDetailRbo)
}

// ReleasePersonDetailRbo 释放PersonDetailRbo
func ReleasePersonDetailRbo(v *PersonDetailRbo) {
	v.PersonTypeList = v.PersonTypeList[:0]
	v.NationalityList = v.NationalityList[:0]
	v.OccupationList = v.OccupationList[:0]
	v.PersonKindList = v.PersonKindList[:0]
	v.MemberList = v.MemberList[:0]
	v.RefShows = v.RefShows[:0]
	v.RefPersons = v.RefPersons[:0]
	v.Name = ""
	v.ThumbUrl = ""
	v.ThumbUrlLg = ""
	v.PosterUrl = ""
	v.PosterUrlH = ""
	v.PersonDesc = ""
	v.Gender = ""
	v.Birthday = ""
	v.Deathday = ""
	v.Homeplace = ""
	v.Height = ""
	v.BloodType = ""
	v.Constellation = ""
	v.Id = 0
	poolPersonDetailRbo.Put(v)
}

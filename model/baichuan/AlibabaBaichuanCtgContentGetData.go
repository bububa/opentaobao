package baichuan

import (
	"sync"
)

// AlibabaBaichuanCtgContentGetData 结构体
type AlibabaBaichuanCtgContentGetData struct {
	// title
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// publishTime
	PublishTime string `json:"publish_time,omitempty" xml:"publish_time,omitempty"`
	// source
	Source string `json:"source,omitempty" xml:"source,omitempty"`
	// summary
	Summary string `json:"summary,omitempty" xml:"summary,omitempty"`
	// coverUrl
	CoverUrl string `json:"cover_url,omitempty" xml:"cover_url,omitempty"`
	// contentUrl
	ContentUrl string `json:"content_url,omitempty" xml:"content_url,omitempty"`
	// paths
	Thumbnails string `json:"thumbnails,omitempty" xml:"thumbnails,omitempty"`
	// org_source
	OrgSource string `json:"org_source,omitempty" xml:"org_source,omitempty"`
	// nick
	Nick string `json:"nick,omitempty" xml:"nick,omitempty"`
}

var poolAlibabaBaichuanCtgContentGetData = sync.Pool{
	New: func() any {
		return new(AlibabaBaichuanCtgContentGetData)
	},
}

// GetAlibabaBaichuanCtgContentGetData() 从对象池中获取AlibabaBaichuanCtgContentGetData
func GetAlibabaBaichuanCtgContentGetData() *AlibabaBaichuanCtgContentGetData {
	return poolAlibabaBaichuanCtgContentGetData.Get().(*AlibabaBaichuanCtgContentGetData)
}

// ReleaseAlibabaBaichuanCtgContentGetData 释放AlibabaBaichuanCtgContentGetData
func ReleaseAlibabaBaichuanCtgContentGetData(v *AlibabaBaichuanCtgContentGetData) {
	v.Title = ""
	v.PublishTime = ""
	v.Source = ""
	v.Summary = ""
	v.CoverUrl = ""
	v.ContentUrl = ""
	v.Thumbnails = ""
	v.OrgSource = ""
	v.Nick = ""
	poolAlibabaBaichuanCtgContentGetData.Put(v)
}

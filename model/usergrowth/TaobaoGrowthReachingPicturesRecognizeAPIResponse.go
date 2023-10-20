package usergrowth

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaogrowthreachingpicturesrecognizeAPIResponse 图片识别 API返回值
// taobao.growth.reaching.pictures.recognize
//
// 图片识别
type TaobaogrowthreachingpicturesrecognizeAPIResponse struct {
	model.CommonResponse
	TaobaogrowthreachingpicturesrecognizeAPIResponseModel
}

// TaobaogrowthreachingpicturesrecognizeAPIResponseModel is 图片识别 成功返回结果
type TaobaogrowthreachingpicturesrecognizeAPIResponseModel struct {
	XMLName xml.Name `xml:"growth_reaching_pictures_recognize_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 识别结果list
	RegionList []RegionData `json:"region_list,omitempty" xml:"region_list>region_data,omitempty"`
	// 图片id
	PicId string `json:"pic_id,omitempty" xml:"pic_id,omitempty"`
	// 唤端链接
	DeepLink string `json:"deep_link,omitempty" xml:"deep_link,omitempty"`
}

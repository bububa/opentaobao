package tuanhotel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripTuanHotelImageUploadAPIResponse 图片上传接口 API返回值
// alitrip.tuan.hotel.image.upload
//
// 用户调用此接口完成外网图片上传至卖家图片中心，此接口返回图片中心的图片地址
type AlitripTuanHotelImageUploadAPIResponse struct {
	model.CommonResponse
	AlitripTuanHotelImageUploadAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripTuanHotelImageUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripTuanHotelImageUploadAPIResponseModel).Reset()
}

// AlitripTuanHotelImageUploadAPIResponseModel is 图片上传接口 成功返回结果
type AlitripTuanHotelImageUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_tuan_hotel_image_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 上传图片结果返回列表
	ImagePathResultList []ImagePathResultVoList `json:"image_path_result_list,omitempty" xml:"image_path_result_list>image_path_result_vo_list,omitempty"`
	// 上传操作错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 上传操作异常信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 上传操作是否成功
	Status bool `json:"status,omitempty" xml:"status,omitempty"`
}

// Reset 清空结构体
func (m *AlitripTuanHotelImageUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ImagePathResultList = m.ImagePathResultList[:0]
	m.MsgCode = ""
	m.Message = ""
	m.Status = false
}

var poolAlitripTuanHotelImageUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripTuanHotelImageUploadAPIResponse)
	},
}

// GetAlitripTuanHotelImageUploadAPIResponse 从 sync.Pool 获取 AlitripTuanHotelImageUploadAPIResponse
func GetAlitripTuanHotelImageUploadAPIResponse() *AlitripTuanHotelImageUploadAPIResponse {
	return poolAlitripTuanHotelImageUploadAPIResponse.Get().(*AlitripTuanHotelImageUploadAPIResponse)
}

// ReleaseAlitripTuanHotelImageUploadAPIResponse 将 AlitripTuanHotelImageUploadAPIResponse 保存到 sync.Pool
func ReleaseAlitripTuanHotelImageUploadAPIResponse(v *AlitripTuanHotelImageUploadAPIResponse) {
	v.Reset()
	poolAlitripTuanHotelImageUploadAPIResponse.Put(v)
}

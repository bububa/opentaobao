package tuanhotel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
图片上传接口 APIResponse
alitrip.tuan.hotel.image.upload

用户调用此接口完成外网图片上传至卖家图片中心，此接口返回图片中心的图片地址
*/
type AlitripTuanHotelImageUploadAPIResponse struct {
    model.CommonResponse
    AlitripTuanHotelImageUploadResponse
}

type AlitripTuanHotelImageUploadResponse struct {
    XMLName xml.Name `xml:"alitrip_tuan_hotel_image_upload_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 上传图片结果返回列表
    
    ImagePathResultList   []ImagePathResultVoList `json:"image_path_result_list,omitempty" xml:"image_path_result_list>image_path_result_vo_list,omitempty"`
    
    
    // 上传操作错误码
    
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`

    
    // 上传操作是否成功
    
    Status   bool `json:"status,omitempty" xml:"status,omitempty"`

    
    // 上传操作异常信息
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
}

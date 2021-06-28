package guoguo

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
上传小件员GPS位置信息 APIResponse
cainiao.guoguo.cp.nborderfrontr.uploadcoordinate

上传小件员GPS位置信息
*/
type CainiaoGuoguoCpNborderfrontrUploadcoordinateAPIResponse struct {
    model.CommonResponse
    CainiaoGuoguoCpNborderfrontrUploadcoordinateResponse
}

type CainiaoGuoguoCpNborderfrontrUploadcoordinateResponse struct {
    XMLName xml.Name `xml:"cainiao_guoguo_cp_nborderfrontr_uploadcoordinate_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 错误信息
    
    StatusMessage   string `json:"status_message,omitempty" xml:"status_message,omitempty"`

    
    // 错误编码
    
    StatusCode   string `json:"status_code,omitempty" xml:"status_code,omitempty"`

    
    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}

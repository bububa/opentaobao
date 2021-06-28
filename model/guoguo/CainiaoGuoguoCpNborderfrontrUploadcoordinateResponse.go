package guoguo

import (
    "github.com/bububa/opentaobao/model"
)

/* 
上传小件员GPS位置信息 APIResponse
cainiao.guoguo.cp.nborderfrontr.uploadcoordinate

上传小件员GPS位置信息
*/
type CainiaoGuoguoCpNborderfrontrUploadcoordinateAPIResponse struct {
    model.CommonResponse
    // Response *CainiaoGuoguoCpNborderfrontrUploadcoordinateResponse `json:"cainiao_guoguo_cp_nborderfrontr_uploadcoordinate_response,omitempty"` 
    CainiaoGuoguoCpNborderfrontrUploadcoordinateResponse
}

/* model for simplify = false
type CainiaoGuoguoCpNborderfrontrUploadcoordinateResponse struct {

    // 错误信息
    
    StatusMessage   string `json:"status_message,omitempty"`
    

    // 错误编码
    
    StatusCode   string `json:"status_code,omitempty"`
    

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

type CainiaoGuoguoCpNborderfrontrUploadcoordinateResponse struct {

    // 错误信息
    StatusMessage   string `json:"status_message,omitempty"`

    // 错误编码
    StatusCode   string `json:"status_code,omitempty"`

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}

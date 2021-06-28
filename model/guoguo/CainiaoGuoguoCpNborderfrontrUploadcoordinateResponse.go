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
	RequestId     string         `json:"request_id,omitempty" xml:"cainiao_guoguo_cp_nborderfrontr_uploadcoordinate_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 错误信息
    
    StatusMessage   string `json:"status_message,omitempty" xml:"
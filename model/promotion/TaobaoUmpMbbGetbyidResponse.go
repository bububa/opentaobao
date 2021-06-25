package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取营销积木块 APIResponse
taobao.ump.mbb.getbyid

根据积木块id获取营销积木块。
*/
type TaobaoUmpMbbGetbyidAPIResponse struct {
    model.CommonResponse
    Response *TaobaoUmpMbbGetbyidResponse `json:"taobao_ump_mbb_getbyid_response,omitempty"`
}

type TaobaoUmpMbbGetbyidResponse struct {

    // 营销积木块定义信息，可以通过ump sdk里面的MBB.fromJson来处理
    Mbb   string `json:"mbb,omitempty"`

}

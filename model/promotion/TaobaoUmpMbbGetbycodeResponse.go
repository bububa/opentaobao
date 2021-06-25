package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
根据营销积木块代码获取积木块 APIResponse
taobao.ump.mbb.getbycode

根据营销积木块代码获取积木块。接口返回该代码最新版本的积木块。如果要查询某个非最新版本的积木块，可以使用积木块id查询接口。
*/
type TaobaoUmpMbbGetbycodeAPIResponse struct {
    model.CommonResponse
    Response *TaobaoUmpMbbGetbycodeResponse `json:"taobao_ump_mbb_getbycode_response,omitempty"`
}

type TaobaoUmpMbbGetbycodeResponse struct {

    // 营销积木块的内容，通过ump sdk来进行处理
    Mbb   string `json:"mbb,omitempty"`

}

package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
删除产品线 APIResponse
taobao.fenxiao.productcat.delete

删除产品线
*/
type TaobaoFenxiaoProductcatDeleteAPIResponse struct {
    model.CommonResponse
    Response *TaobaoFenxiaoProductcatDeleteResponse `json:"taobao_fenxiao_productcat_delete_response,omitempty"`
}

type TaobaoFenxiaoProductcatDeleteResponse struct {

    // 操作是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}

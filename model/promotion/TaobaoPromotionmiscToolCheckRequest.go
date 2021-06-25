package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
UMP工具检测 APIRequest
taobao.promotionmisc.tool.check

UMP工具检测。ISV通过该接口检测（通过taobao.ump.tool.add）创建的UMP工具（tool）是否符合规范，如果不符合，则返回错误信息和对应的解决方案的；工具检测通过后才可以提交工具审核邮件，提交工具审核时，需提供该接口的返回值。
*/
type TaobaoPromotionmiscToolCheckRequest struct {
    model.Params

    // 工具ID, taobao.ump.tool.add成功后返回的id。
    toolId   int64 

    // 可使用的元数据。PRD审核后，会告诉isv可使用的元数据。
    metaAllow   string 

}

func NewTaobaoPromotionmiscToolCheckRequest() *TaobaoPromotionmiscToolCheckRequest{
    return &TaobaoPromotionmiscToolCheckRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPromotionmiscToolCheckRequest) GetApiMethodName() string {
    return "taobao.promotionmisc.tool.check"
}

func (r TaobaoPromotionmiscToolCheckRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPromotionmiscToolCheckRequest) SetToolId(toolId int64) error {
    r.toolId = toolId
    r.Set("tool_id", toolId)
    return nil
}

func (r TaobaoPromotionmiscToolCheckRequest) GetToolId() int64 {
    return r.toolId
}

func (r *TaobaoPromotionmiscToolCheckRequest) SetMetaAllow(metaAllow string) error {
    r.metaAllow = metaAllow
    r.Set("meta_allow", metaAllow)
    return nil
}

func (r TaobaoPromotionmiscToolCheckRequest) GetMetaAllow() string {
    return r.metaAllow
}


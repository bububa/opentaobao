package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取分销商信息 API返回值 
taobao.fenxiao.distributors.get

查询和当前登录供应商有合作关系的分销商的信息
*/
type TaobaoFenxiaoDistributorsGetAPIResponse struct {
    model.CommonResponse
    TaobaoFenxiaoDistributorsGetResponse
}

// 获取分销商信息 成功返回结果
type TaobaoFenxiaoDistributorsGetResponse struct {
    XMLName xml.Name `xml:"fenxiao_distributors_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 分销商详细信息
    Distributors   []Distributor `json:"distributors,omitempty" xml:"distributors>distributor,omitempty"`
}

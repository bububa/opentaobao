package wlb

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询家装服务商列表 API返回值 
taobao.wlb.order.jzpartner.query

为支持家装类目的商家，对绑定家装物流服务的订单可以在商家的ERP中发货、批量发货，因此开发根据服务类型查询所有的服务商列表的接口
*/
type TaobaoWlbOrderJzpartnerQueryAPIResponse struct {
    model.CommonResponse
    TaobaoWlbOrderJzpartnerQueryResponse
}

// 查询家装服务商列表 成功返回结果
type TaobaoWlbOrderJzpartnerQueryResponse struct {
    XMLName xml.Name `xml:"wlb_order_jzpartner_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口查询成功或者失败
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
    // 查询返回信息，如果失败，存储错误信息
    ResultInfo   string `json:"result_info,omitempty" xml:"result_info,omitempty"`
    // 物流配送服务商对象列表
    ServerList   []PartnerNew `json:"server_list,omitempty" xml:"server_list>partner_new,omitempty"`
    // 安装服务商列表
    InstallList   []PartnerNew `json:"install_list,omitempty" xml:"install_list>partner_new,omitempty"`
}

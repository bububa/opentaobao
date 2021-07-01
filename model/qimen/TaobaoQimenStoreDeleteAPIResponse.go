package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
门店删除接口 API返回值 
taobao.qimen.store.delete

商家在ERP等系统中调用该接口，删除线下门店
*/
type TaobaoQimenStoreDeleteAPIResponse struct {
    model.CommonResponse
    TaobaoQimenStoreDeleteAPIResponseModel
}

// 门店删除接口 成功返回结果
type TaobaoQimenStoreDeleteAPIResponseModel struct {
    XMLName xml.Name `xml:"qimen_store_delete_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 响应信息
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // 响应标示
    Flag   string `json:"flag,omitempty" xml:"flag,omitempty"`
    // 响应code
    QimenCode   string `json:"qimen_code,omitempty" xml:"qimen_code,omitempty"`
}

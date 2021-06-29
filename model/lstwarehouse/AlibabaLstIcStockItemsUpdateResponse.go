package lstwarehouse

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
零售通经销商商品库存设置 APIResponse
alibaba.lst.ic.stock.items.update

零售通经销商商品库存设置
*/
type AlibabaLstIcStockItemsUpdateAPIResponse struct {
    model.CommonResponse
    AlibabaLstIcStockItemsUpdateResponse
}

type AlibabaLstIcStockItemsUpdateResponse struct {
    XMLName xml.Name `xml:"alibaba_lst_ic_stock_items_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaLstIcStockItemsUpdateResult `json:"result,omitempty" xml:"result,omitempty"`

    
}

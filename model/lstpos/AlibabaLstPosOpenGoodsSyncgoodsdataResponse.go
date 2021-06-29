package lstpos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
门店商品批量同步接口(最多10条商品信息) APIResponse
alibaba.lst.pos.open.goods.syncgoodsdata

门店商品批量同步接口(最多10条商品信息)
*/
type AlibabaLstPosOpenGoodsSyncgoodsdataAPIResponse struct {
    model.CommonResponse
    AlibabaLstPosOpenGoodsSyncgoodsdataResponse
}

type AlibabaLstPosOpenGoodsSyncgoodsdataResponse struct {
    XMLName xml.Name `xml:"alibaba_lst_pos_open_goods_syncgoodsdata_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果对象
    
    Result   *AlibabaLstPosOpenGoodsSyncgoodsdataResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}

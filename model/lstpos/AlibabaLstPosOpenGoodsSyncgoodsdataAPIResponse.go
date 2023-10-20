package lstpos

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstPosOpenGoodsSyncgoodsdataAPIResponse 门店商品批量同步接口(最多10条商品信息) API返回值
// alibaba.lst.pos.open.goods.syncgoodsdata
//
// 门店商品批量同步接口(最多10条商品信息)
type AlibabaLstPosOpenGoodsSyncgoodsdataAPIResponse struct {
	model.CommonResponse
	AlibabaLstPosOpenGoodsSyncgoodsdataAPIResponseModel
}

// AlibabaLstPosOpenGoodsSyncgoodsdataAPIResponseModel is 门店商品批量同步接口(最多10条商品信息) 成功返回结果
type AlibabaLstPosOpenGoodsSyncgoodsdataAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_pos_open_goods_syncgoodsdata_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果对象
	Result *AlibabaLstPosOpenGoodsSyncgoodsdataResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

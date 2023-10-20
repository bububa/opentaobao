package lstpos

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *AlibabaLstPosOpenGoodsSyncgoodsdataAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLstPosOpenGoodsSyncgoodsdataAPIResponseModel).Reset()
}

// AlibabaLstPosOpenGoodsSyncgoodsdataAPIResponseModel is 门店商品批量同步接口(最多10条商品信息) 成功返回结果
type AlibabaLstPosOpenGoodsSyncgoodsdataAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_pos_open_goods_syncgoodsdata_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果对象
	Result *AlibabaLstPosOpenGoodsSyncgoodsdataResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLstPosOpenGoodsSyncgoodsdataAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaLstPosOpenGoodsSyncgoodsdataAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLstPosOpenGoodsSyncgoodsdataAPIResponse)
	},
}

// GetAlibabaLstPosOpenGoodsSyncgoodsdataAPIResponse 从 sync.Pool 获取 AlibabaLstPosOpenGoodsSyncgoodsdataAPIResponse
func GetAlibabaLstPosOpenGoodsSyncgoodsdataAPIResponse() *AlibabaLstPosOpenGoodsSyncgoodsdataAPIResponse {
	return poolAlibabaLstPosOpenGoodsSyncgoodsdataAPIResponse.Get().(*AlibabaLstPosOpenGoodsSyncgoodsdataAPIResponse)
}

// ReleaseAlibabaLstPosOpenGoodsSyncgoodsdataAPIResponse 将 AlibabaLstPosOpenGoodsSyncgoodsdataAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLstPosOpenGoodsSyncgoodsdataAPIResponse(v *AlibabaLstPosOpenGoodsSyncgoodsdataAPIResponse) {
	v.Reset()
	poolAlibabaLstPosOpenGoodsSyncgoodsdataAPIResponse.Put(v)
}

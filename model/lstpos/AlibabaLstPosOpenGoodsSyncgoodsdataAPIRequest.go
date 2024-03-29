package lstpos

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstPosOpenGoodsSyncgoodsdataAPIRequest 门店商品批量同步接口(最多10条商品信息) API请求
// alibaba.lst.pos.open.goods.syncgoodsdata
//
// 门店商品批量同步接口(最多10条商品信息)
type AlibabaLstPosOpenGoodsSyncgoodsdataAPIRequest struct {
	model.Params
	// 商品对象列表
	_goodsDTOList []GoodsDto
	// 用户主账号Id
	_userId int64
}

// NewAlibabaLstPosOpenGoodsSyncgoodsdataRequest 初始化AlibabaLstPosOpenGoodsSyncgoodsdataAPIRequest对象
func NewAlibabaLstPosOpenGoodsSyncgoodsdataRequest() *AlibabaLstPosOpenGoodsSyncgoodsdataAPIRequest {
	return &AlibabaLstPosOpenGoodsSyncgoodsdataAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLstPosOpenGoodsSyncgoodsdataAPIRequest) Reset() {
	r._goodsDTOList = r._goodsDTOList[:0]
	r._userId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstPosOpenGoodsSyncgoodsdataAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.pos.open.goods.syncgoodsdata"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstPosOpenGoodsSyncgoodsdataAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLstPosOpenGoodsSyncgoodsdataAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGoodsDTOList is GoodsDTOList Setter
// 商品对象列表
func (r *AlibabaLstPosOpenGoodsSyncgoodsdataAPIRequest) SetGoodsDTOList(_goodsDTOList []GoodsDto) error {
	r._goodsDTOList = _goodsDTOList
	r.Set("goods_d_t_o_list", _goodsDTOList)
	return nil
}

// GetGoodsDTOList GoodsDTOList Getter
func (r AlibabaLstPosOpenGoodsSyncgoodsdataAPIRequest) GetGoodsDTOList() []GoodsDto {
	return r._goodsDTOList
}

// SetUserId is UserId Setter
// 用户主账号Id
func (r *AlibabaLstPosOpenGoodsSyncgoodsdataAPIRequest) SetUserId(_userId int64) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabaLstPosOpenGoodsSyncgoodsdataAPIRequest) GetUserId() int64 {
	return r._userId
}

var poolAlibabaLstPosOpenGoodsSyncgoodsdataAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLstPosOpenGoodsSyncgoodsdataRequest()
	},
}

// GetAlibabaLstPosOpenGoodsSyncgoodsdataRequest 从 sync.Pool 获取 AlibabaLstPosOpenGoodsSyncgoodsdataAPIRequest
func GetAlibabaLstPosOpenGoodsSyncgoodsdataAPIRequest() *AlibabaLstPosOpenGoodsSyncgoodsdataAPIRequest {
	return poolAlibabaLstPosOpenGoodsSyncgoodsdataAPIRequest.Get().(*AlibabaLstPosOpenGoodsSyncgoodsdataAPIRequest)
}

// ReleaseAlibabaLstPosOpenGoodsSyncgoodsdataAPIRequest 将 AlibabaLstPosOpenGoodsSyncgoodsdataAPIRequest 放入 sync.Pool
func ReleaseAlibabaLstPosOpenGoodsSyncgoodsdataAPIRequest(v *AlibabaLstPosOpenGoodsSyncgoodsdataAPIRequest) {
	v.Reset()
	poolAlibabaLstPosOpenGoodsSyncgoodsdataAPIRequest.Put(v)
}

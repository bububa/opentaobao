package lstpos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstPosOpenGoodsGetgoodsbypagingAPIRequest 分页查询用户全量的门店域商品接口(每页最多20条) API请求
// alibaba.lst.pos.open.goods.getgoodsbypaging
//
// 分页查询用户全量的门店域商品接口(每页最多20条)
type AlibabaLstPosOpenGoodsGetgoodsbypagingAPIRequest struct {
	model.Params
	// 当前页码
	_page int64
	// 当前主账号userId
	_userId int64
}

// NewAlibabaLstPosOpenGoodsGetgoodsbypagingRequest 初始化AlibabaLstPosOpenGoodsGetgoodsbypagingAPIRequest对象
func NewAlibabaLstPosOpenGoodsGetgoodsbypagingRequest() *AlibabaLstPosOpenGoodsGetgoodsbypagingAPIRequest {
	return &AlibabaLstPosOpenGoodsGetgoodsbypagingAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstPosOpenGoodsGetgoodsbypagingAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.pos.open.goods.getgoodsbypaging"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstPosOpenGoodsGetgoodsbypagingAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetPage is Page Setter
// 当前页码
func (r *AlibabaLstPosOpenGoodsGetgoodsbypagingAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r AlibabaLstPosOpenGoodsGetgoodsbypagingAPIRequest) GetPage() int64 {
	return r._page
}

// SetUserId is UserId Setter
// 当前主账号userId
func (r *AlibabaLstPosOpenGoodsGetgoodsbypagingAPIRequest) SetUserId(_userId int64) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabaLstPosOpenGoodsGetgoodsbypagingAPIRequest) GetUserId() int64 {
	return r._userId
}

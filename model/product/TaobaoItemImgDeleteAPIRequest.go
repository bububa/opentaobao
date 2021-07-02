package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItemImgDeleteAPIRequest 删除商品图片 API请求
// taobao.item.img.delete
//
// 删除商品图片
type TaobaoItemImgDeleteAPIRequest struct {
	model.Params
	// 商品数字ID
	_numIid int64
	// 商品图片ID；如果是竖图，请将id的值设置为1
	_id int64
	// 标记是否要删除第6张图，因为第6张图与普通商品图片不是存储在同一个位置的无图片ID，所以要通过一个标记来判断是否为第6张图，目前第6张图业务主要用在女装业务下
	_isSixthPic bool
}

// NewTaobaoItemImgDeleteRequest 初始化TaobaoItemImgDeleteAPIRequest对象
func NewTaobaoItemImgDeleteRequest() *TaobaoItemImgDeleteAPIRequest {
	return &TaobaoItemImgDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoItemImgDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.item.img.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoItemImgDeleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is NumIid Setter
// 商品数字ID
func (r *TaobaoItemImgDeleteAPIRequest) SetNumIid(_numIid int64) error {
	r._numIid = _numIid
	r.Set("num_iid", _numIid)
	return nil
}

// Get NumIid Getter
func (r TaobaoItemImgDeleteAPIRequest) GetNumIid() int64 {
	return r._numIid
}

// Set is Id Setter
// 商品图片ID；如果是竖图，请将id的值设置为1
func (r *TaobaoItemImgDeleteAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// Get Id Getter
func (r TaobaoItemImgDeleteAPIRequest) GetId() int64 {
	return r._id
}

// Set is IsSixthPic Setter
// 标记是否要删除第6张图，因为第6张图与普通商品图片不是存储在同一个位置的无图片ID，所以要通过一个标记来判断是否为第6张图，目前第6张图业务主要用在女装业务下
func (r *TaobaoItemImgDeleteAPIRequest) SetIsSixthPic(_isSixthPic bool) error {
	r._isSixthPic = _isSixthPic
	r.Set("is_sixth_pic", _isSixthPic)
	return nil
}

// Get IsSixthPic Getter
func (r TaobaoItemImgDeleteAPIRequest) GetIsSixthPic() bool {
	return r._isSixthPic
}

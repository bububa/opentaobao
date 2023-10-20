package tbitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoitemimgdeleteAPIRequest 删除商品图片 API请求
// taobao.item.img.delete
//
// 删除商品图片
type TaobaoitemimgdeleteAPIRequest struct {
	model.Params
	// 商品数字ID
	_numIid int64
	// 商品图片ID；如果是竖图，请将id的值设置为1
	_id int64
	// 标记是否要删除第6张图，因为第6张图与普通商品图片不是存储在同一个位置的无图片ID，所以要通过一个标记来判断是否为第6张图，目前第6张图业务主要用在女装业务下
	_isSixthPic bool
}

// NewTaobaoitemimgdeleteRequest 初始化TaobaoitemimgdeleteAPIRequest对象
func NewTaobaoitemimgdeleteRequest() *TaobaoitemimgdeleteAPIRequest {
	return &TaobaoitemimgdeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoitemimgdeleteAPIRequest) GetApiMethodName() string {
	return "taobao.item.img.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoitemimgdeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoitemimgdeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNumIid is NumIid Setter
// 商品数字ID
func (r *TaobaoitemimgdeleteAPIRequest) SetNumIid(_numIid int64) error {
	r._numIid = _numIid
	r.Set("num_iid", _numIid)
	return nil
}

// GetNumIid NumIid Getter
func (r TaobaoitemimgdeleteAPIRequest) GetNumIid() int64 {
	return r._numIid
}

// SetId is Id Setter
// 商品图片ID；如果是竖图，请将id的值设置为1
func (r *TaobaoitemimgdeleteAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r TaobaoitemimgdeleteAPIRequest) GetId() int64 {
	return r._id
}

// SetIsSixthPic is IsSixthPic Setter
// 标记是否要删除第6张图，因为第6张图与普通商品图片不是存储在同一个位置的无图片ID，所以要通过一个标记来判断是否为第6张图，目前第6张图业务主要用在女装业务下
func (r *TaobaoitemimgdeleteAPIRequest) SetIsSixthPic(_isSixthPic bool) error {
	r._isSixthPic = _isSixthPic
	r.Set("is_sixth_pic", _isSixthPic)
	return nil
}

// GetIsSixthPic IsSixthPic Getter
func (r TaobaoitemimgdeleteAPIRequest) GetIsSixthPic() bool {
	return r._isSixthPic
}

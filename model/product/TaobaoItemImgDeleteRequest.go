package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除商品图片 APIRequest
taobao.item.img.delete

删除商品图片
*/
type TaobaoItemImgDeleteRequest struct {
    model.Params

    // 商品数字ID
    numIid   int64 

    // 商品图片ID；如果是竖图，请将id的值设置为1
    id   int64 

    // 标记是否要删除第6张图，因为第6张图与普通商品图片不是存储在同一个位置的无图片ID，所以要通过一个标记来判断是否为第6张图，目前第6张图业务主要用在女装业务下
    isSixthPic   bool 

}

func NewTaobaoItemImgDeleteRequest() *TaobaoItemImgDeleteRequest{
    return &TaobaoItemImgDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoItemImgDeleteRequest) GetApiMethodName() string {
    return "taobao.item.img.delete"
}

func (r TaobaoItemImgDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoItemImgDeleteRequest) SetNumIid(numIid int64) error {
    r.numIid = numIid
    r.Set("num_iid", numIid)
    return nil
}

func (r TaobaoItemImgDeleteRequest) GetNumIid() int64 {
    return r.numIid
}

func (r *TaobaoItemImgDeleteRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

func (r TaobaoItemImgDeleteRequest) GetId() int64 {
    return r.id
}

func (r *TaobaoItemImgDeleteRequest) SetIsSixthPic(isSixthPic bool) error {
    r.isSixthPic = isSixthPic
    r.Set("is_sixth_pic", isSixthPic)
    return nil
}

func (r TaobaoItemImgDeleteRequest) GetIsSixthPic() bool {
    return r.isSixthPic
}


package fenxiao

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/fenxiao"
)

/* 
添加产品 
taobao.fenxiao.product.add

添加分销平台产品数据。业务逻辑与分销系统前台页面一致。<br/><br/>    * 产品图片默认为空<br/>    * 产品发布后默认为下架状态
*/
func TaobaoFenxiaoProductAdd(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoProductAddRequest, session string) (*fenxiao.TaobaoFenxiaoProductAddAPIResponse, error) {
    var resp fenxiao.TaobaoFenxiaoProductAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

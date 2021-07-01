package fenxiao

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/fenxiao"
)

/* 
分销商等级查询 
taobao.fenxiao.grades.get

根据供应商ID，查询他的分销商等级信息
*/
func TaobaoFenxiaoGradesGet(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoGradesGetAPIRequest, session string) (*fenxiao.TaobaoFenxiaoGradesGetAPIResponse, error) {
    var resp fenxiao.TaobaoFenxiaoGradesGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

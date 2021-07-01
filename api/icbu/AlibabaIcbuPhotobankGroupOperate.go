package icbu

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/icbu"
)

/* 
图片银行分组操作接口 
alibaba.icbu.photobank.group.operate

修改用户图片银行的分组信息，包括 新增分组，删除分组，分组重命名
*/
func AlibabaIcbuPhotobankGroupOperate(clt *core.SDKClient, req *icbu.AlibabaIcbuPhotobankGroupOperateAPIRequest, session string) (*icbu.AlibabaIcbuPhotobankGroupOperateAPIResponse, error) {
    var resp icbu.AlibabaIcbuPhotobankGroupOperateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}

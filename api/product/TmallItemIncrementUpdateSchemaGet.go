package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
天猫增量更新商品规则获取 
tmall.item.increment.update.schema.get

增量方式修改天猫商品的规则获取的API。<br/>1.接口返回支持增量修改的字段以及相应字段的规则。<br/>2.如果入参xml_data指定了更新的字段，则只返回指定字段的规则（ISV如果功能性很强，如明确更新Title，请拼装好次字段以提升API整体性能）；<br/>3.ISV初次接入，开发阶段，此字段不填可以看到所有支持增量的字段；但是如果上线功能明确，请尽量遵守第2条<br/>4.如果ISV对字段规则非常清晰，可以直接组装入参数据提交到tmall.item.schema.increment.update进行数据更新。但是最好不要写死，比如每天还是有对此接口功能的一次比对。<br/>---（感谢爱慕旗舰店提供API命名）
*/
func TmallItemIncrementUpdateSchemaGet(clt *core.SDKClient, req *product.TmallItemIncrementUpdateSchemaGetRequest, session string) (*product.TmallItemIncrementUpdateSchemaGetResponse, error) {
    var resp product.TmallItemIncrementUpdateSchemaGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}

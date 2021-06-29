package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
实体动态更新 APIRequest
alibaba.ailabs.aligenie.skill.entity.import

根据用户上传的实体信息，进行制定技能实体的动态变更
*/
type AlibabaAilabsAligenieSkillEntityImportRequest struct {
    model.Params

    // 技能Id
    skillId   int64 

    // 是否增量更新，true为增量，否则上传数据中没有的实体值将被删除
    append   bool 

    // 测试，目前没有此功能，此参数无效
    test   bool 

    // 文件内容
    fileData   []*model.File 

    // 要更新的实体名
    entityName   string 

    // 固定值，ISV
    type   string 

}

func NewAlibabaAilabsAligenieSkillEntityImportRequest() *AlibabaAilabsAligenieSkillEntityImportRequest{
    return &AlibabaAilabsAligenieSkillEntityImportRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAilabsAligenieSkillEntityImportRequest) GetApiMethodName() string {
    return "alibaba.ailabs.aligenie.skill.entity.import"
}

func (r AlibabaAilabsAligenieSkillEntityImportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAilabsAligenieSkillEntityImportRequest) SetSkillId(skillId int64) error {
    r.skillId = skillId
    r.Set("skill_id", skillId)
    return nil
}

func (r AlibabaAilabsAligenieSkillEntityImportRequest) GetSkillId() int64 {
    return r.skillId
}

func (r *AlibabaAilabsAligenieSkillEntityImportRequest) SetAppend(append bool) error {
    r.append = append
    r.Set("append", append)
    return nil
}

func (r AlibabaAilabsAligenieSkillEntityImportRequest) GetAppend() bool {
    return r.append
}

func (r *AlibabaAilabsAligenieSkillEntityImportRequest) SetTest(test bool) error {
    r.test = test
    r.Set("test", test)
    return nil
}

func (r AlibabaAilabsAligenieSkillEntityImportRequest) GetTest() bool {
    return r.test
}

func (r *AlibabaAilabsAligenieSkillEntityImportRequest) SetFileData(fileData []*model.File) error {
    r.fileData = fileData
    r.Set("file_data", fileData)
    return nil
}

func (r AlibabaAilabsAligenieSkillEntityImportRequest) GetFileData() []*model.File {
    return r.fileData
}

func (r *AlibabaAilabsAligenieSkillEntityImportRequest) SetEntityName(entityName string) error {
    r.entityName = entityName
    r.Set("entity_name", entityName)
    return nil
}

func (r AlibabaAilabsAligenieSkillEntityImportRequest) GetEntityName() string {
    return r.entityName
}

func (r *AlibabaAilabsAligenieSkillEntityImportRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r AlibabaAilabsAligenieSkillEntityImportRequest) GetType() string {
    return r.type
}


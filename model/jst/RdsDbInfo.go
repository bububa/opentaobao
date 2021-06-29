package jst

// RdsDbInfo 
type RdsDbInfo struct {
    // 最大帐号数，1个数据库最多可以创建的账户数目
    MaxAccount   string `json:"max_account,omitempty" xml:"max_account,omitempty"`
    // 数据库登录密码
    Password   string `json:"password,omitempty" xml:"password,omitempty"`
    // 数据库状态 0：创建中 ；1：激活；3：正在删除
    DbStatus   string `json:"db_status,omitempty" xml:"db_status,omitempty"`
    // 数据库编码
    Charset   string `json:"charset,omitempty" xml:"charset,omitempty"`
    // rds实例id
    InstanceId   string `json:"instance_id,omitempty" xml:"instance_id,omitempty"`
    // 登录数据库的帐号
    UserName   string `json:"user_name,omitempty" xml:"user_name,omitempty"`
    // 数据库名称
    DbName   string `json:"db_name,omitempty" xml:"db_name,omitempty"`
    // 用户id
    Uid   string `json:"uid,omitempty" xml:"uid,omitempty"`
    // rds实例名
    InstanceName   string `json:"instance_name,omitempty" xml:"instance_name,omitempty"`
    // rds实例类型,s:共享型，x:专享型
    InstanceType   string `json:"instance_type,omitempty" xml:"instance_type,omitempty"`
    // 数据库类型，mysql或者mssql
    DbType   string `json:"db_type,omitempty" xml:"db_type,omitempty"`
    // 备注
    Comment   string `json:"comment,omitempty" xml:"comment,omitempty"`
}

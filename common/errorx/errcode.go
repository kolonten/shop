package errorx

// OK 成功返回
const OK uint32 = 200

/**(前3位代表业务,后三位代表具体功能)**/

// 全局错误码

//REQUEST_ERROR 请求错误
const REQUEST_ERROR uint32 = 100400

//REQUEST_PARAM_ERROR 请求参数错误
const REQUEST_PARAM_ERROR uint32 = 100002

//TOKEN_EXPIRE_ERROR 令牌过期错误
const TOKEN_EXPIRE_ERROR uint32 = 100401

//TOKEN_GENERATE_ERROR 令牌生成错误
const TOKEN_GENERATE_ERROR uint32 = 100403

// SERVER_COMMON_ERROR 服务器常见错误
const SERVER_COMMON_ERROR uint32 = 100500

//DB_ERROR 数据库错误
const DB_ERROR uint32 = 100520

//DB_UPDATE_AFFECTED_ZERO_ERROR 数据库更新影响零错误
const DB_UPDATE_AFFECTED_ZERO_ERROR uint32 = 100521

//用户模块

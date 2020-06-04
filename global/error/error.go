package error

const (
	// 系统部分
	ErrorsConfigInitFail         string = "初始化配置文件发生错误"
	ErrorsFuncEventAlreadyExists string = "注册函数类事件失败，键名已经被注册"
	ErrorsFuncEventNotRegister   string = "没有找到键名对应的函数"
	ErrorsFuncEventNotCall       string = "注册的函数无法正确执行"
	ErrorsBasePath               string = "Init base path error"
	ErrorsNoAuthorization        string = "token鉴权未通过，请通过token授权接口重新获取token,"
	// 数据库部分
	ErrorsDbDriverNotExists   string = "数据库驱动不存在"
	ErrorsDbSqlDriverInitFail string = "数据库驱动初始化失败"
	ErrorsDbGetConnFail       string = "从数据库连接池获取一个连接失败，超过最大连接次数."
	ErrorsDbPrepareRunFail    string = "sql语句预处理（prepare）失败"
	ErrorsDbQueryRunFail      string = "查询类sql语句执行失败"
	ErrorsDbExecuteRunFail    string = "执行类sql语句执行失败"
	// redis部分
	ErrorsRedisInitConnFail string = "初始化redis连接池失败"
	ErrorsRedisAuhtFail     string = "Redis Auht鉴权失败，密码错误"
	// 验证器错误
	ErrorsValiadatorNotExists string = "不存在的验证器"
	// token部分
	ErrorsTokenInvalid string = "无效的token"
	ErrorsTokenExpired string = "过期的token"
	// snowflake
	ErrorsSnowflakeInitFail  string = "初始化 snowflakeFctory 过程发生错误"
	ErrorsSnowflakeGetIdFail string = "获取snowflake唯一ID过程发生错误"
	// websocket
	ErrorsWebsocketOnOpenFail string = "websocket onopen 发生阶段错误"
	// rabbitMq
	ErrorsRabbitMqReconnectFail string = "RabbitMq消费者端掉线后重连失败，超过尝试最大次数"
)

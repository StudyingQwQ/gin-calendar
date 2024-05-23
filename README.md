# 基于gin+mysql+redis+gorm+viper的日程提醒demo

## 启动
1. 修改配置文件`config.yaml`
2. 创建数据库 `gin_calendar`
3. 执行命令 `go run main.go`即可自动建表

## 内容
1. 用户登陆注册
2. redis缓存日程，到达时间后会提醒
3. 日程的增删改查
4. 邮箱提醒(需要有有能收得到邮件的电子邮箱)
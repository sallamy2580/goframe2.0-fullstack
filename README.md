# hotgo


#### HotGo is a full-stack development base platform and mobile application platform based on vue and goframe2.0 with front-end and back-end separation. It integrates jwt authentication, dynamic routing, dynamic menu, casbin authentication, message queue, timed tasks and other functions. A variety of common scene files, allowing you to focus more time on business development

## Technical selection

* Backend: Use goframe2.0 to quickly build basic API, goframe2.0 is a web framework written in go language.
* Front-end: Build basic pages based on JeeSite Mobile Uni-App+aidex-sharp.
* Database: Using MySql (8.0) version, use gorm to implement basic operations on the database.
* Cache: Use Redis to record the jwt token of the current active user and implement multi-sign-on restrictions.
* API Documentation: Use Swagger to build automated documentation.
* Message queue: Compatible with kafka, redis, rocketmq at the same time, one-click configuration to switch to the MQ you want to use.

## System screenshot
#### * web side

![image](https://user-images.githubusercontent.com/26652343/155689571-e6a0a5a3-011b-44cc-b84b-a1c82301b207.png)

![image](https://user-images.githubusercontent.com/26652343/155689646-d3395261-6061-469f-8256-3cd0ff9f5d05.png)

![image](https://user-images.githubusercontent.com/26652343/155689709-5ddac1d3-1c01-4fab-9d3a-9ece72ca5ba0.png)

#### * mobile
![image](https://user-images.githubusercontent.com/26652343/155689481-2fc019eb-18e4-4a94-b417-50524e945089.png)
![image](https://user-images.githubusercontent.com/26652343/155689738-ac97f9c0-47ae-499b-b3fe-0cb4ce97f3bc.png)

## Environmental requirements
- node version >= v14.0.0 
- golang version >= v1.16
- IDEversion：Goland
- mysqlversion >=8.0
- redisversion >=5.0

## quick start
1. Pull the code to the server where you have installed the above environment
 ```shell script
git clone https://github.com/bufanyun/hotgo.git
 ```

2. Configure your site information

Server:
 - Create a mysql database, import the database file into your mysql, directory address: /hotgo-server/storage/hotgo.sql
 - Change /hotgo-server/config/config.example.yaml to: config.yaml, and configure it according to your actual environment

web+uinapp side:
 - Configure the server address, which is included in the following file:
 * hotgo-uniapp/common/config.js 
 * /hotgo-uniapp/manifest.json 
 * hotgo-uniapp/common/config.js 

3. Start the service
Server:
   ```shell script
  cd hotgo-server
  go mod tidy  #update package
  go run main.go  #start the service
```

web side:
   ```shell script
cd hotgo-web
npm install #Install dependencies
npm run dev #start web project
```
uinapp side:
- 1、Download and install: Integrated Development Environment HBuilderX (recommended, VSCode or WebStorm can also be used)
- 2、Menu: File -> Import -> Import from local directory, select the "jeesite4-uniapp" folder.
- 3、Menu: Run -> Run to Built-in Browser (or Run to Browser -> Chrome Browser).
- 4、After the HBuliderX console is compiled, the mobile phone login page will pop up automatically.


## Special thanks to(The following ranks are in no particular order)

* goframe2.0 https://goframe.org
* JeeSite Mobile Uni-App https://gitee.com/thinkgem/jeesite4-uniapp
* aidex-sharp https://gitee.com/big-hedgehog/aidex-sharp

## Open Source Statement
* At present, the project is still being updated continuously and is only for reference and learning. If you encounter any problems, please contact the author on the WeChat below!

![image](https://user-images.githubusercontent.com/26652343/155691271-1ded98d8-f0f1-4467-9079-26cec1195af5.png)
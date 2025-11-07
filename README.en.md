# Smart Work Order System

The Smart-API Work Order System is a backend management system developed in Go, featuring a front-end and back-end separation. The back-end uses the Gin framework, while the front-end uses Vue.js and Element UI. This system aims to efficiently handle work order management, including task allocation, execution monitoring, and real-time updates.

<img align="right" width="320" src="https://github.com/sunwenbo/golang/raw/master/logo.png">

[![Build Status](https://github.com/sunwenbo/smart/workflows/build/badge.svg)](https://github.com/sunwenbo/smart-api)
[![Static Badge](https://img.shields.io/badge/release-v1.0.0-blue)](https://github.com/sunwenbo/smart-api/releases)
[![License](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/sunwenbo/smart-api)

简体中文 | [English](https://github.com/sunwenbo/smart-api/blob/main/README.Zh-cn.md)

Based on Gin + Vue + Element UI (or Arco Design/Ant Design), the Smart-API system is a front-end and back-end separated work order management system. Initialization is simple, with just minor database connection configuration needed. The system supports multiple command operations, including database migration commands for simplified initialization.

## SMART Naming Principles
- **S**: Specific - Clear and well-defined goals, not ambiguous or vague.
- **M**: Measurable - Quantifiable metrics to evaluate progress.
- **A**: Achievable - Realistic and feasible to accomplish.
- **R**: Relevant - Aligned with work priorities and able to drive outcomes.
- **T**: Time-bound - Defined timeframe to ensure completion within deadlines.


[Frontend Project](https://github.com/sunwenbo/smart-ui)

[Backend Project](https://github.com/sunwenbo/smart-api)

## 🎬 Online Demo

[Click to access online demo](http://117.50.46.228/login)

> Default login: `admin` / `123456`

## ✨ Architectural Features

- Follows RESTful API design standards
- Built on the Gin framework with middleware support (user authentication, CORS, access logs, trace IDs, etc.)
- Casbin-based RBAC access control model
- JWT authentication
- Swagger documentation support (via swaggo)
- GORM-based database storage, with support for multiple databases
- Simple config file mapping for quick setup
- Code generation tool
- Form building tool
- Multi-command mode
- Multi-tenancy support
- TODO: Unit tests

## 🤩 Key Features

- **Custom Work Order Forms**: Create custom forms for various work order types.
- **Custom Approval Processes**: Flexible assignment of approvers and multi-person collaborative approval.
- **LDAP Login Support**: Compatible with enterprise internal LDAP authentication.
- **Work Order Favorites**: Multiple users can save frequently used work orders.
- **Work Order Management**: Comprehensive management of task creation, assignment, tracking, and execution status.
- **Work Order Reminders**: Send notification messages, including integration with third-party platforms (e.g., Feishu, DingTalk).
- **Real-Time Task Monitoring**: Real-time updates of task execution status using WebSocket.
- **Role-Based Access Control**: Fine-grained role and permission management based on Casbin.
- **Ratings and Comments**: After task completion, users can rate and comment on the work order.
- **Log Recording**: Detailed operation and task execution logs for auditing and analysis.
- **Charts and Stats**: Visualized work order statistics, including weekly and monthly reports, and personal submission rankings.

## 🎁 Built-In Modules

1. **Multi-Tenancy**: Supports multi-tenancy by separating databases for each tenant.
2. **User Management**: Configure system user settings.
3. **Department Management**: Organize company departments and teams.
4. **Role Management**: Manage role-based access control and data permissions.
5. **Menu Management**: Configure system menus, button permissions, and API access.
6. **Dictionary Management**: Manage commonly used fixed data in the system.
7. **Parameter Management**: Dynamically configure system parameters.
8. **Operation Logs**: Record and query system operation logs and error logs.
9. **Login Logs**: Record and query login logs, including abnormal login events.
10. **API Documentation**: Automatically generate API documentation based on business logic.
11. **Code Generation**: Visualized code generation based on database tables for CRUD operations.
12. **Form Building**: Customizable drag-and-drop page layout for forms.
13. **Service Monitoring**: View basic server information.
14. **Content Management**: Demo content management module for quick reference.
15. **Scheduled Tasks**: Automate task execution, supporting both API and function calls.

## 🛰️ System Modules

1. **User Management**: Manage user information and permission assignment.
2. **Task Center**: Manage work order assignment, tracking, and monitoring.
3. **Work Order Center**: Submit and view work orders.
4. **Process Center**: Customize work order types, forms, and approval workflows.
5. **Role Management**: Fine-grained permission control via roles.
6. **Log Management**: Manage system operation and task execution logs.
7. **Service Monitoring**: Real-time monitoring of server performance and status.
8. **Ratings and Comments**: Provide ratings and feedback for completed tasks.

## 🔧 Technology Stack

- **Backend**: Go, Gin, GORM, JWT, Casbin
- **Frontend**: Vue.js, Element UI, Axios, WebSocket
- **Database**: MySQL / PostgreSQL / SQLite
- **Tools**: Docker, Swagger, Git

## 🚀 Getting Started

### Prerequisites

You need to have [Go](https://go.dev/dl/), [Gin](https://pkg.go.dev/github.com/gin-gonic/gin), [Node.js](http://nodejs.org/), and [Git](https://git-scm.com/) installed locally.

### 📦 Environment Requirements

- Go 1.18 or later
- Node.js v14.16.0 or later
- npm version 6.14.11 or later
- MySQL 8.0 or another compatible database
- Docker, Kubernetes (optional)

### Backend Installation

1. Clone the project:

    ```bash
    git clone https://github.com/sunwenbo/smart-api.git
    cd smart-api
    ```

2. Install dependencies:

    ```bash
    go mod tidy
    ```

3. Configure database connection:

   ```bash
   cp config/settings.full.yml config/settings.yml
   ```
   ⚠️ : Modify the database connection information in config/settings.yml to ensure that the database configuration is correct. Create the database manually before you initialize it.
   1. Modify the database information in the configuration file
   2. Note the configuration data in settings.database
   3. Confirm the log path

4. Compile
    ```bash 
    go build -o smart-api main.go
    ```
5. Initialize the Database:

   ```bash
   # The first configuration requires initializing the database resource information.
   # Use on macOS or Linux
   $ ./smart-api migrate -c config/settings.yml
   
   # ⚠️ Note: Use on Windows
   $ smart-api.exe migrate -c config/settings.yml
   
   # Start the project; you can also debug it using an IDE.
   # Use on macOS or Linux
   $ ./smart-api server -c config/settings.yml
   
   # ⚠️ Note: Use on Windows
   $ smart-api.exe server -c config/settings.yml

6. How to Add Data to the `sys_api` Table:

   ```bash
   # When the project starts, using `-a true` will automatically add missing interface data.
   ./smart-api server -c config/settings.yml -a true

7. Update Swagger Documentation Generation:

   ```bash
   go generate
   ```

8. Start Service：

    ```bash
    ./smart-api server -c config/settings.yml
    ```
9. Local Development Database Migration:

   ```
   Generate migration files:
   go run main.go migrate -g true -c config/settings.yml
   
   Generate migration files -- system-related:
   go run main.go migrate -g true -a true -c config/settings.yml
   
   After modifying the migration files, execute the command below to start changes:
   go run main.go migrate -c config/settings.yml
   
   Interface management generation:
   go run main.go server -c config/settings.yml -a false
   ```
#### Cross-Compilation

```bash
# windows
env GOOS=windows GOARCH=amd64 go build main.go

# or
# linux
env GOOS=linux GOARCH=amd64 go build main.go
```

### Frontend Installation Steps

1. Clone the frontend project:

    ```bash
    git clone https://github.com/sunwenbo/smart-ui.git
    cd smart-ui
    ```

2. Install dependencies:

    ```bash
    npm config set registry https://registry.npmmirror.com/  
    npm config get registry
    npm install --legacy-peer-deps 
    ```

3. Start the development server:

    ```bash
    npm run dev
    ```

4. Access the address:

   Open your browser and visit `http://localhost:9527` to view the frontend interface.


### Start with Docker
If not deploying locally, you need to initialize the database data first. You can use the locally compiled binary file to execute the following command.

⚠️⚠️⚠️ Modify the database information in the `settings.yml` file.
```bash 
./smart-api migrate -c config/settings.yml 
```

#### Backend
If you are on an ARM architecture locally but want to build an x86 architecture image:
```bash
# Build the image
docker buildx build --platform linux/amd64 -t registry.cn-beijing.aliyuncs.com/sunwenbo/smart-api:latest . --load    
# Push to the image repository
docker push registry.cn-beijing.aliyuncs.com/sunwenbo/smart-api:latest
# Start the backend service, mount the settings.yml file
docker run -itd  \
  -p 8000:8000 \
  -v /data/config/:/home/service/config/ \
  --name smart-api \
  registry.cn-beijing.aliyuncs.com/sunwenbo/smart-api:latest   
```

#### Frontend
```bash
# Build the image
docker buildx build --platform linux/amd64 -t registry.cn-beijing.aliyuncs.com/sunwenbo/smart-ui:latest . --load
# Push to the image repository
docker push registry.cn-beijing.aliyuncs.com/sunwenbo/smart-ui:latest 
# Start the frontend service, mount the Nginx configuration file
docker run -itd  \
  -p 80:80 \
  -v /data/default.conf:/etc/nginx/conf.d/default.conf \
  --name smart-ui \
  registry.cn-beijing.aliyuncs.com/sunwenbo/smart-ui:latest  
```

####  View the post-deployment status
![img.png](static%2Fimage%2Fimg.png)

### Deploy with Kubernetes
The method to initialize data is as mentioned above ☝️

#### Backend
```bash
cd smart-api/scripts/
bash install.sh
```

#### Frontend
```bash
cd smart-ui/scripts/
# Before executing deployment, modify the configurations in the yaml file, such as ingress
bash install.sh
```

#### Check Deployment Status
```bash
kubectl get cm -n sre 
kubectl get pod -n sre 
kubectl get ingress -n sre 
```

## Usage Instructions

1. **Log into the system**: Use the default admin account `admin` to log in.
2. **Create a work order**: Go to the task management page, click "Create Work Order", fill in the relevant information and assign a handler.
3. **Task execution monitoring**: Through the system's task execution interface, you can view the execution status, results, and logs in real time.
4. **Work order rating and comments**: After the work order is completed, users can rate the work order and use the comment function to record feedback on the processing of the work order.


## System Screenshots
[Data Statistics]![Data Statistics](https://github.com/sunwenbo/smart-api/blob/main/static/image/statistics.png)

[Work Order Application]![Work Order Application](https://github.com/sunwenbo/smart-api/blob/main/static/image/orderApply-v2.png)

[Custom Process]![Custom Process](https://github.com/sunwenbo/smart-api/blob/main/static/image/customflow.png)

[Custom Template]![Custom Template](https://github.com/sunwenbo/smart-api/blob/main/static/image/customtemplate.png)

[Work Order Management]![Work Order Management](https://github.com/sunwenbo/smart-api/blob/main/static/image/ordermanage.png)

[Task Type]![Task Type](https://github.com/sunwenbo/smart-api/blob/main/static/image/tasktype.png)

### Login Interface
![Login Interface](https://github.com/sunwenbo/smart-api/blob/main/static/image/login.png)

### Work Order Management
![Work Order Management](https://github.com/sunwenbo/smart-api/blob/main/static/image/order.png)

### Real-time Task Monitoring
![Real-time Task Monitoring](https://github.com/sunwenbo/smart-api/blob/main/static/image/task.png)

## Contribution Guidelines

We welcome community developers to contribute code or provide feedback. If you would like to participate in this project, please follow these steps:

1. **Fork this repository**: Click the `Fork` button in the upper right corner to copy the project to your GitHub account.
2. **Clone your fork**: Run the following command in the terminal to clone the project to your local machine:

    ```bash
    git clone https://github.com/your-username/smart-api-backend.git
    ```

3. **Create a branch**: In the cloned project directory, create a new branch for development:

    ```bash
    git checkout -b feature-branch-name
    ```
4. **Commit your changes**: After development, use `git add` and `git commit` to submit your code changes.

    ```bash
    git add .
    git commit -m "Describe your changes"
    ```

5. **Submit a Pull Request**: Submit a Pull Request on GitHub, ensuring your code passes all tests and adheres to coding standards.

## FAQ

### How to modify database configurations?

Database configurations are stored in `config/settings.yml`. You can modify the `database` configuration items according to your needs.

### How to add new API interfaces?

Create new handling logic in the `api` directory and register that route in the `router`. The service layer logic should be placed in the `service` directory.

### How to implement WebSocket for real-time task monitoring?

During task execution, the system pushes task status updates to the frontend via WebSocket. The frontend uses the `WebSocket` interface to receive and dynamically display task execution logs.

## Open Source License

The Smart-API work order system is open-sourced under the [MIT License](LICENSE). Personal and commercial use and modifications are welcome.


## 📨 Interaction

<table>
   <tr>
    <td><img src="https://github.com/sunwenbo/golang/raw/master/wx.jpeg" width="180px"></td>
    <td><img src="https://github.com/sunwenbo/golang/raw/master/wx.jpeg" width="180px"></td>
</tr>
  <tr>
    <td>Personal WeChat</td>
    <td>WeChat Group 🔥🔥🔥</td>
  </tr>

</table>

## 💎 Contributors

Thanks to the following developers for their contributions to this project:

- [Developer 1](https://github.com/developer1)
- [Developer 2](https://github.com/developer2)
- [Developer 3](https://github.com/developer3)


> If you have any questions or suggestions, please submit an [issue](https://github.com/sunwenbo/smart-api/issues).


## 🤝 Special Thanks

### JetBrains Open Source License Support

The `smart-api` project has always been developed using JetBrains' GoLand IDE, based on **free JetBrains Open Source license(s)**. Thank you for your support.

<a href="https://www.jetbrains.com/?from=kubeadm-ha" target="_blank"><img src="https://raw.githubusercontent.com/panjf2000/illustrations/master/jetbrains/jetbrains-variant-4.png" width="250" align="middle"/></a>

### Open Source Framework Components
1. [go-admin # 非常强大的后端开发框架、内置功能非常丰富](https://github.com/go-admin-team/go-admin)
2. [wfd-vue # 流程设计器](https://github.com/guozhaolong/wfd-vue)
3. [VForm 一款高效的Vue 2 / Vue3 的低代码表单可视化设计，一键生成源码](https://vform666.com/vform3.html)
4. [ant-design](https://github.com/ant-design/ant-design)
5. [ant-design-pro](https://github.com/ant-design/ant-design-pro)
6. [arco-design](https://github.com/arco-design/arco-design)
7. [arco-design-pro](https://github.com/arco-design/arco-design-pro)
8. [gin](https://github.com/gin-gonic/gin)
9. [casbin](https://github.com/casbin/casbin)
10. [spf13/viper](https://github.com/spf13/viper)
11. [gorm](https://github.com/jinzhu/gorm)
12. [gin-swagger](https://github.com/swaggo/gin-swagger)
13. [jwt-go](https://github.com/dgrijalva/jwt-go)
14. [vue-element-admin](https://github.com/PanJiaChen/vue-element-admin)
15. [ruoyi-vue](https://gitee.com/y_project/RuoYi-Vue)
16. [form-generator](https://github.com/JakHuang/form-generator)

## 🤟 Donation

> If you find this project helpful, you can buy the author a cup of ☕️ as encouragement:

<img class="no-margin" src="https://github.com/sunwenbo/golang/raw/master/wxPay.jpeg" height="200px" >

## 🤝 Links

[Go Developer Growth Roadmap](http://www.golangroadmap.com/)

## 🔑 License

[MIT](https://github.com/sunwenbo/smart-api/blob/main/LICENSE.md)

Copyright (c) 2022 sunwenbo

# Study-fiber
This repository is about some demos I learned about the fiber framework, it's not for production use, there's a bunch of things to do if you want to improve.

# User module

```yaml
path: /api/v1/users/:id
method: GET
description: 获取{id}的信息 # get user information by id
```

```yaml
path: /api/v1/users/:id
method: PUT
description: 提交{id}的信息 # submiting user information by id
```

```yaml
path: /api/v1/users/:id
method: DELETE
description: 删除指定id的用户信息(只标记为删除,在后续的开发中系统自动删除 (保留90天))
```


# Post module

* post API
```yaml
path: /api/v1/post
method: POST
description: "创建一个帖子" # create a post
```
* response
```json
{
  "status": 200,
  "data":[
    {
      "id":120,
      "context":""
    }
  ]
}
```



```yaml
path: /api/v1/post/:id
method: DELETE
description: 删除指定id的帖子 # delete post by id
```

```yaml
path: /api/v1/post/:id
method: GET
description: 获取指定id的帖子 # get post by id
```

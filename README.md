# User module

```yaml
path: /api/v1/users/:id
method: GET
description: 获取{id}的信息
```

```yaml
path: /api/v1/users/:id
method: PUT
description: 提交{id}的信息
```

```yaml
path: /api/v1/users/:id
method: DELETE
description: 删除指定id的用户信息(只标记为删除,在后续的开发中系统自动删除 (保留90天))
```


# Post module

```yaml
path: /api/v1/post
method: POST
description: "创建一个帖子"
```
* 返回数据
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
description: 删除指定id的帖子
```

```yaml
path: /api/v1/post/:id
method: GET
description: 获取指定id的帖子
```

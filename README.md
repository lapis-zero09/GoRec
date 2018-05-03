# GoRec


## Methods

- User base
    - Pearson Coef

### input data


```
user_id, item_id, rating
1, 2, 4
1, 127, 1
21, 63, 2
...
```


### Example

```sh
$ cd gorec/src/data
$ sh ./download_data.sh
```

```sh
$ cd gorec/src
$ go run main.go
```

<img src="img/sim.png">


---

change main.go

```go
~~~
MostSimilarUser(encountered, userItemMatrix, userId, similarSize)
~~~
```

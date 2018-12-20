# GoRec

blog: https://lapis-zero09.hatenablog.com/entry/2018/05/06/171536

## Methods

- User/Item base
    - Pearson
    - Cosine
    - Adjusted Cosine
    - Jaccard
    - Dice
    - Simpson

## input data

```
user_id, item_id, rating
1, 2, 4
1, 127, 1
21, 63, 2
...
```
ex) http://files.grouplens.org/papers/ml-100k.zip

<img src="img/data.png">


## Example

http://files.grouplens.org/papers/ml-100k.zip

### Download Data

```sh
$ cd gorec
$ sh ./scripts/download_data.sh
```

### Calc simiralities
```sh
$ cd gorec
$ go run cmd/gorec/main.go ./data/ml-100k/u.data
```

### Output

```
»»»» go run cmd/gorec/main.go ./data/ml-100k/u.data
2018/12/21 04:11:59 UserSize: 943
2018/12/21 04:11:59 ItemSize: 1682

+++User Simlarity+++

Adjusted Cosine
-----------------------------
2018/12/21 04:12:18 target ID: 941
rank    id       similarity
-----------------------------
 1       91      0.243949
 2       780     0.242358
-----------------------------
2018/12/21 04:12:18 target ID: 356
rank    id       similarity
-----------------------------
 1       101     0.462062
 2       518     0.446907
 3       181     0.430992
-----------------------------


+++Item Simlarity+++

Pearson
-----------------------------
2018/12/21 04:12:38 target ID: 941
rank    id       similarity
-----------------------------
 1       91      0.373956
 2       454     0.370661
-----------------------------
2018/12/21 04:12:38 target ID: 356
rank    id       similarity
-----------------------------
 1       518     0.480340
 2       181     0.477340
 3       101     0.496726
-----------------------------

Cosine
-----------------------------
2018/12/21 04:12:42 target ID: 941
rank    id       similarity
-----------------------------
 1       454     0.414486
 2       91      0.405048
-----------------------------
2018/12/21 04:12:42 target ID: 356
rank    id       similarity
-----------------------------
 1       101     0.516649
 2       181     0.510658
 3       518     0.501872
-----------------------------
```


### change main.go

You can change some parameter for getting similar user/item.
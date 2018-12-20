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
<img src="img/sim.png">


### change main.go

You can change some parameter for getting similar user/item.
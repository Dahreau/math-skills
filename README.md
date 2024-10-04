# math-skills

## Objectives

The project is meant to calculate and display the following:
- the [**Average**](https://en.wikipedia.org/wiki/Average)[<sub> FR</sub>](https://fr.wikipedia.org/wiki/Moyenne_arithm%C3%A9tique)
- the [**Median**](https://en.wikipedia.org/wiki/Median)[<sub> FR</sub>](https://fr.wikipedia.org/wiki/M%C3%A9diane_(statistiques))
- the [**Variance**](https://en.wikipedia.org/wiki/Variance)[<sub> FR</sub>](https://fr.wikipedia.org/wiki/Variance_(math%C3%A9matiques))
- the [**Standart Deviation**](https://en.wikipedia.org/wiki/Standard_deviation)[<sub> FR</sub>](https://fr.wikipedia.org/wiki/%C3%89cart_type)

Note : Each value is rounded as to not display any decimals.

## Installation

Firstly, you can clone the git repository with this command :
```
https://zone01normandie.org/git/dsamaky/math-skills
```

Secondly, the program needs a list of random integers (unsorted or sorted) to work. The list can be represented in two different ways :

- separated by **spaces** :

```
1 2 3
```

- separated by **line breaks** :
```
1
2
3
```

## Usage

To run the project with the list being named _data.txt_, the user has to type the following command :

```
go run . data.txt
```


## Result

As an exemple, the content of _data.txt_ will be :

```
1 2 3 4 5
```

The program should display the following result :
```
Average: 3
Median: 3
Variance: 2
Standart Deviation: 1
```
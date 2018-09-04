# acolor
![GitHub top language](https://img.shields.io/github/languages/top/schdav/acolor.svg)
![license](https://img.shields.io/github/license/schdav/acolor.svg)
![GitHub last commit](https://img.shields.io/github/last-commit/schdav/acolor.svg)
![GitHub repo size in bytes](https://img.shields.io/github/repo-size/schdav/acolor.svg)

acolor computes the average color of a jpeg image.

## How to use?
`go build`

`./acolor IMAGE`

## Why exponentiations and square roots?
These mathematical operations are needed to calculate the most accurate average color.  
(See [Computer Color is Broken](https://www.youtube.com/watch?v=LKnqECcg6Gw))

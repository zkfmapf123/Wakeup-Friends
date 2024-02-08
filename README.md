# Wake-up-Friends

## Desc
- AWS Console 들어가서 EC2 인스턴스를 Running / Stopped 하기 귀찮음
- 그냥 보고 깨우자 (Running)

## Required

- AWS Configure default 로 설정이 되있어야 한다.
- Region 명시 되어있어야 한다.

## Exec 

```sh
    ## use go file
    make dev

    ## use bin
    make run
```

## 구성

```
    make build
    mv ./bin/main ./bin/friends
    mv /usr/local/bin ./bin/friends
    friends
```

## 기능

- Dashboard

![dashboard](./public/dashboardgif.gif)

- Wakeup

![wakeup](./public/wakeup.gif)

- Sleep

![sleep](./public/sleep.gif)
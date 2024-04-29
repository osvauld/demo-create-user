### How to run on server

```
screen -S demo_create_user
cd demo-create-user
go run .
```


### Screen commands

# attach the screen

```
screen -r demo_create_user
```

# detach and attach
```
screen -rd demo_create_user
```

# list screens
```
screen -ls
```


# show current screen
```
echo $STY
```
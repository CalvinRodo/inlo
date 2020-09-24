# Incident Log CLI 

## Usage

Start an incident

```
inlo start Everything is on fire
```

Assign an Incident Commander 
```
inlo ic Calvin Rodo 
```

Add to timeline
```
inlo tl this is a new item
```

Add an action itme 
```
inlo ai we should do this
```

Attach a file to the log 
```
inlo attach ./filename.yaml
```

Attach the output from a command to the log 
```
whois alpha.canada.ca | inlo read whois_result
```

End an incident
```
inlo end
```


## Configuration

*LogDir* The location to store the incident logs and attached files
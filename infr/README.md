```bash
$ docker-compose up -d
$ docker exec -it mng_cont bash
...

# mongo -u herodotus -p herodotus

> use herodotus
switched to db herodotus
> db.campaigns.insert({"name": "some"})
```
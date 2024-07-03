
Проверить, что все подсети в сети `default` отсортированы совпадают с ожидаемым списком:
```bash
yc-query << 'EOF'
sort(map(Cloud("cloud").Folder("default").Network("default").Subnets(), {.Name})) == ["default-ru-central1-a", "default-ru-central1-b", "default-ru-central1-d"]
EOF
```

Получить список всех фолдеров в облаке:
```bash
yc-query <<'EOF'
Cloud("cloud").Folders() | toJSON()
EOF
```

Получить список всех инстансов в облаке и найти для каждого инстанса данные о его boot диске:
```bash
yc-query << 'EOF'
map(Cloud("cloud").Folder("default").Instances(), .BootDisk()) | toJSON()
EOF
```

Получить список всех фолдеров в облаке и для каждого фолдера найти список сетей. Вывести в формате JSON:
```bash
yc-query << 'EOF'
map(Cloud("cloud").Folders(), {let x = {"name":.Name,  "nets":.Networks()}; x}) | toJSON()
EOF
```

Найти в облаке `cloud` фолдер `default`, в нем найти инстанс `reverse-vpn` и получить список его дисков.
Далее отфильтровать диски, размер которых больше 10Гб и посчитать общий объем таких дисков:
```bash
yc-query << 'EOF'
let disks = Cloud("cloud").Folder("default").Instance("reverse-vpn").Disks();
let big_disks = filter(disks, .Size > 10 * 1024 * 1024 * 1024);
FormatSize(sum(map(big_disks, .Size)))
EOF
```

В облаке `cloud` найти все сервисные аккаунты, созданные за последние 2 года и отсортировать их по дате создания: 
```bash
yc-query <<'EOF'
let folder = Cloud("cloud").Folder("default");
let accounts = folder.ServiceAccounts();

accounts | filter(.CreatedAt > now() - Duration("P2Y")) | sortBy(.CreatedAt, "desc") | toJSON()
EOF
```


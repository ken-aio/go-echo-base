# What's this

goのWeb Frameworkであるechoを使う際のコピー元のbaseとして使えるプロジェクトです  
参考のプロジェクトをベースにしてテストコードの追加をしました  

looperを使って自動テストをwatchすると開発がやりやすいです  
https://github.com/nathany/looper

DBのmigrationにはRubyのActiveRecordを使います  

```
$ cd migration
$ bundle install
$ bundle exec rake db:create RAKE_ENV=development
$ bundle exec rake db:migrate RAKE_ENV=development
```

参考  
https://github.com/eurie-inc/echo-sample

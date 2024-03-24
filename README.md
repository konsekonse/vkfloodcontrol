В файле main.go необходимо реализовать интерфейс FloodControl, который будет проверять флуд-контроль для вызовов метода Check. При этом требуется иметь возможность обрабатывать запросы из нескольких экземпляров приложения, поэтому необходимо использовать общее хранилище данных.

Мой первый шаг заключался в определении структуры данных для хранения информации о вызовах метода Check для каждого пользователя и выборе подходящего механизма синхронизации

Я рассмотрел несколько вариантов реализации хранилища данных:

Использование глобальной переменной: В этом случае можно было бы использовать глобальную переменную для хранения информации о вызовах метода Check для каждого пользователя. Но такой подход не является безопасным с точки зрения параллелизма, поскольку доступ к глобальным переменным несинхронизирован в многопоточной среде, что может привести к состоянию гонки и непредсказуемому поведению программы.

Использование мьютекса и мапы: Вместо этого я выбрал использование мьютекса для синхронизации доступа к карте, которая будет отображать идентификаторы пользователей на времена их вызовов метода Check. Этот подход гарантирует безопасность доступа к данным и избегает дата рейса.

Использование каналов: Еще один вариант был использовать каналы для передачи информации о вызовах метода Check между различными горутинами. Однако в этом случае сложнее было бы управлять сроками хранения информации и удалять устаревшие данные из канала.

Я решил выбрать второй вариант, используя мьютекс и мапу, поскольку он предоставляет простой и надежный механизм для безопасного доступа к данным. Этот подход также обеспечивает гибкость в управлении временем хранения для каждого пользователя
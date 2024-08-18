# Dependency Injection 복습을 위한 예제

## Dependency Injection 이란?
모듈간의 Dependency 해소를 위해 `Dependnecy Injector`를 사용하여 모듈간의 의존성을 해소하는 방법을 말함. 
이해에 도움이 되었던 [stack overflow 답변](https://stackoverflow.com/questions/130794/what-is-dependency-injection/44945310#44945310)

## 현 레포의 예시를 자세히 풀어보자
- 현재 예시는 `Dependency Injection` 중에서도 `Constructor Injection`을 사용한 예시이다.
  - `Constructor Injection`은 생성자를 통해 의존성을 주입하는 방법이다.
- `Constructor Injection`을 사용하면 `Dependency`를 주입받는 객체가 생성될 때 `Dependency`를 주입받아 사용할 수 있다.

- `Developer` 인터페이스는 `Developer`를 구현한 클래스들이 구현해야 하는 메소드를 정의한다.
- `BackendDeveloper`와 `FrontendDeveloper`는 `Developer` 인터페이스를 구현한 클래스이다.
- `Project` 클래스는 `Developer` 인터페이스를 구현한 클래스를 멤버변수로 가지고 있다.
- `Project` 클래스의 생성자에서 `Developer` 인터페이스를 구현한 클래스를 주입받아 멤버변수로 사용한다.
- `Project` 클래스의 `startProject` 메소드를 호출하면 `Developer` 인터페이스를 구현한 클래스의 `develop` 메소드를 호출한다.
- `Main` 클래스에서 `Project` 클래스의 객체를 생성할 때 `Developer` 인터페이스를 구현한 클래스를 주입하여 사용한다.
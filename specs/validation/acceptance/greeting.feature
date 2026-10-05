Feature: Greeting

  @story-1
  Rule: A named greeting addresses the caller by name

    Scenario: A client requests a greeting with a name
      Given the greeter service is running
      When a Client calls "GET /hello" with name "Ada"
      Then the response is a JSON greeting that addresses "Ada"

  @story-2
  Rule: A greeting without a name still succeeds

    Scenario: A client requests a greeting with no name
      Given the greeter service is running
      When a Client calls "GET /hello" with no name
      Then the response is a successful JSON greeting with a default message

  @story-3
  Rule: A named farewell addresses the caller by name

    Scenario: A client requests a farewell with a name
      Given the greeter service is running
      When a Client calls "GET /farewell" with name "Ada"
      Then the response is a JSON goodbye message that addresses "Ada"

  @story-4
  Rule: A farewell without a name still succeeds

    Scenario: A client requests a farewell with no name
      Given the greeter service is running
      When a Client calls "GET /farewell" with no name
      Then the response is a successful JSON goodbye message with a default message

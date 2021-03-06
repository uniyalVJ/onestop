package org.cedar.onestop.user

import org.cedar.onestop.user.config.AuthorizationConfiguration
import org.hamcrest.Matchers
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.boot.test.context.SpringBootTest
import org.springframework.http.MediaType
import org.springframework.security.test.context.support.WithMockUser
import org.springframework.security.test.web.servlet.setup.SecurityMockMvcConfigurers
import org.springframework.test.context.ActiveProfiles
import org.springframework.test.web.servlet.MockMvc
import org.springframework.test.web.servlet.ResultActions
import org.springframework.test.web.servlet.request.MockMvcRequestBuilders
import org.springframework.test.web.servlet.result.MockMvcResultMatchers
import org.springframework.test.web.servlet.setup.MockMvcBuilders
import org.springframework.web.context.WebApplicationContext
import spock.lang.Specification

import static org.springframework.boot.test.context.SpringBootTest.WebEnvironment.RANDOM_PORT

@ActiveProfiles(['integration'])
@SpringBootTest(classes = [UserApplication.class], webEnvironment = RANDOM_PORT)
class SavedSearchControllerIntegrationSpec extends Specification {

  @Autowired
  private WebApplicationContext context

  private MockMvc mvc

  def setup() {
    mvc = MockMvcBuilders
        .webAppContextSetup(context)
        .apply(SecurityMockMvcConfigurers.springSecurity())
        .build()

    mvc.perform(MockMvcRequestBuilders
        .post("/v1/user")
        .contentType("application/json")
        .content(('{ "id":"mockMvcUser", "name": "test"}'))
        .accept(MediaType.APPLICATION_JSON))
  }

  def "unauthenticated user is NOT authorized and gets translated by controller advice"() {
    when: 'We make a request to a endpoint beyond our scope'
    def getResults = mvc.perform(MockMvcRequestBuilders
        .get("/v1/saved-search")
        .accept(MediaType.APPLICATION_JSON))

    then: 'we get the translated controller advice response'
    getResults.andExpect(MockMvcResultMatchers.status().isUnauthorized())
  }

  @WithMockUser(username = "mockMvcUser", roles = AuthorizationConfiguration.LIST_ALL_SAVED_SEARCHES)
  def "admin user authorized to admin getAll endpoint"() {
    when: 'We make a request to a endpoint beyond our scope'
    def getResults = mvc.perform(MockMvcRequestBuilders
        .get("/v1/saved-search")
        .accept(MediaType.APPLICATION_JSON))

    then: 'we get 200'
    getResults.andExpect(MockMvcResultMatchers.status().isOk())
  }

  @WithMockUser(username = "mockMvcUser", roles = [AuthorizationConfiguration.READ_OWN_PROFILE])
  def "public user cannot get saved searches belonging to other users"() {
    setup: 'must have a user to associate the search with'
    mvc.perform(MockMvcRequestBuilders
        .post("/v1/self")
        .contentType("application/json")
        .content(('{"id":"mockMvcUser"}'))
        .accept(MediaType.APPLICATION_JSON))

    when: 'We make a request to a endpoint beyond our scope'
    def getResults = mvc.perform(MockMvcRequestBuilders
        .get("/v1/saved-search/user/{id}", "unknownUser")
        .accept(MediaType.APPLICATION_JSON))

    then: 'we get the translated controller advice response'
    getResults.andExpect(MockMvcResultMatchers.status().isNotFound())
  }

  @WithMockUser(username = "mockMvcUser", roles = [AuthorizationConfiguration.READ_OWN_PROFILE, AuthorizationConfiguration.CREATE_SAVED_SEARCH, AuthorizationConfiguration.READ_SAVED_SEARCH])
  def "POST and GET save search items, user is taken from Authentication principal"() {
    setup: 'must have a user to associate the search with'
    mvc.perform(MockMvcRequestBuilders
        .post("/v1/self")
        .contentType("application/json")
        .content(('{ "id":"mockMvcUser"}'))
        .accept(MediaType.APPLICATION_JSON))

    when:
    ResultActions postOneResults = mvc.perform(MockMvcRequestBuilders
        .post("/v1/self/saved-search")
        .contentType("application/json")
        .content('{ "name": "testOne", "value": "valueOne" }')
        .accept(MediaType.APPLICATION_JSON))

    ResultActions postTwoResults = mvc.perform(MockMvcRequestBuilders
        .post("/v1/self/saved-search")
        .contentType("application/json")
        .content('{ "name": "testTwo", "value": "valueTwo" }')
        .accept(MediaType.APPLICATION_JSON))

    then:
    postOneResults.andExpect(MockMvcResultMatchers.status().isCreated())
        .andExpect(MockMvcResultMatchers.content().contentType(MediaType.APPLICATION_JSON))
        .andExpect(MockMvcResultMatchers.jsonPath("\$.data[0].attributes.name").value("testOne"))
        .andExpect(MockMvcResultMatchers.jsonPath("\$.data[0].attributes.value").value("valueOne"))

    postTwoResults.andExpect(MockMvcResultMatchers.status().isCreated())
        .andExpect(MockMvcResultMatchers.content().contentType(MediaType.APPLICATION_JSON))
        .andExpect(MockMvcResultMatchers.jsonPath("\$.data[0].attributes.name").value("testTwo"))
        .andExpect(MockMvcResultMatchers.jsonPath("\$.data[0].attributes.value").value("valueTwo"))

    when:
    def getResults = mvc.perform(MockMvcRequestBuilders
        .get("/v1/self/saved-search")
        .accept(MediaType.APPLICATION_JSON))

    then:
    getResults.andExpect(MockMvcResultMatchers.status().isOk())
        .andExpect(MockMvcResultMatchers.content().contentType(MediaType.APPLICATION_JSON))
        .andExpect(MockMvcResultMatchers.jsonPath("\$.data.length()").value(2))
        .andExpect((MockMvcResultMatchers.jsonPath("\$.data[*].attributes.name", Matchers.containsInAnyOrder("testOne", "testTwo"))))
        .andExpect((MockMvcResultMatchers.jsonPath("\$.data[*].attributes.value", Matchers.containsInAnyOrder("valueOne", "valueTwo"))))
  }

}

package org.cedar.psi.registry.api

import groovy.json.JsonSlurper
import groovy.util.logging.Slf4j
import org.apache.kafka.streams.KafkaStreams
import org.apache.kafka.streams.state.QueryableStoreTypes
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.beans.factory.annotation.Value
import org.springframework.web.bind.annotation.PathVariable
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RestController

import javax.servlet.http.HttpServletResponse

import static org.springframework.web.bind.annotation.RequestMethod.GET
import static org.springframework.web.bind.annotation.RequestMethod.HEAD

@Slf4j
@RestController
class MetadataRestController {

  @Value('${kafka.stores.raw.granule}')
  String RAW_GRANULE_STORE

  @Value('${kafka.stores.raw.collection}')
  String RAW_COLLECTION_STORE

  private KafkaStreams metadataStream

  @Autowired
  MetadataRestController(KafkaStreams metadataStream) {
    this.metadataStream = metadataStream
  }


  @RequestMapping(path = '/metadata/{type}/{id}', method = [GET, HEAD], produces = 'application/json')
  Map retrieveJson(@PathVariable String type, @PathVariable String id, HttpServletResponse response) {
    def storeName =
        type == 'granule' ? RAW_GRANULE_STORE :
            type == 'collection' ? RAW_COLLECTION_STORE : null

    if (storeName) {
      def value = getFromStreamStore(metadataStream, storeName, id)
      if (value) {
        return [id: id, value: value]
      }
    }

    response.sendError(404)
    return null
  }

  private static Map getFromStreamStore(KafkaStreams streamsApp, String storeName, String id) {
    def store = streamsApp.store(storeName, QueryableStoreTypes.keyValueStore())
    return store ? new JsonSlurper().parseText(store.get(id) as String) as Map : null
  }


//  @RequestMapping(path='/metadata/rebuild', method = [PUT], produces = 'application/json')
//  Map rebuild() {
//    this.metadataStream.close()
//    this.metadataStream.cleanUp()
//    this.metadataStream.start()
//    return ['acknowledged': true]
//  }

}

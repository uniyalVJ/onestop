import Immutable from 'seamless-immutable'
import {
  COLLECTION_SEARCH_START,
  COLLECTION_SEARCH_COMPLETE,
  COLLECTION_SEARCH_ERROR,
} from '../../actions/search/CollectionRequestActions'

export const initialState = Immutable({
  collectionSearchRequestInFlight: false,
})

export const collectionRequest = (state = initialState, action) => {
  switch (action.type) {
    case COLLECTION_SEARCH_START:
      return Immutable.set(state, 'collectionSearchRequestInFlight', true)

    case COLLECTION_SEARCH_COMPLETE:
      return Immutable.set(state, 'collectionSearchRequestInFlight', false)

    case COLLECTION_SEARCH_ERROR:
      return Immutable.set(state, 'collectionSearchRequestInFlight', false)

    default:
      return state
  }
}

export default collectionRequest

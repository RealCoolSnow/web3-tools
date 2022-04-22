import { BrowserRouter, Switch, Route } from 'react-router-dom'
import routes from '@/router'

const App = () => {
  return (
    <BrowserRouter>
      <Switch>
        {routes.map((route) => (
          <Route exact key={route.path} path={route.path}>
            <route.component />
          </Route>
        ))}
      </Switch>
    </BrowserRouter>
  )
}

export default App

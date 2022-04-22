import { BrowserRouter, Routes, Route } from 'react-router-dom'
import routes from '@/router'

const App = () => {
  return (
    <BrowserRouter>
      <Routes>
        {routes.map((route) => (
          <Route exact key={route.path} path={route.path}>
            <route.component />
          </Route>
        ))}
      </Routes>
    </BrowserRouter>
  )
}

export default App

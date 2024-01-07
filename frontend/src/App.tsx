import "./App.css";
import Login from "./pages/Login";
import { Route, BrowserRouter as Router, Routes } from "react-router-dom";
import TwoFactorAuth from "./pages/TwoFactorAuth";

function App() {
  return (
    <>
      <Router>
        <Routes>
          <Route path="" element={<Login />} />
          <Route path="/challenge" element={<TwoFactorAuth/>}/>
          
        </Routes>
      </Router>
    </>
  );
}

export default App;

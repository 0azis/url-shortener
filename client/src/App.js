import './App.css';
import {SignIn} from "./components/SignIn/signin";
import {BrowserRouter, Route, Routes} from "react-router-dom";
import {NavBar} from "./components/NavBar/nav";
import {SignUp} from "./components/SignUp/signup";
import {Home} from "./components/Home/home";

function App() {
  return (
    <BrowserRouter>
    <Routes>
        <Route path="*" element={<SignIn />} />
        <Route path="/signup" element={<SignUp />} />
        <Route path="/home" element={[<NavBar />, <Home />]} />
    </Routes>
    </BrowserRouter>
  );
}

export default App;

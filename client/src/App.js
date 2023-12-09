import './App.css';
import {SignIn} from "./components/SignIn/signin";
import {BrowserRouter, Route, Routes} from "react-router-dom";
import {NavBar} from "./components/NavBar/nav";
import {SignUp} from "./components/SignUp/signup";
import {Home} from "./components/Home/home";
import {Url} from "./components/Url/url";

function App() {
  return (
    <BrowserRouter>
    <Routes>
        <Route path="*" element={<SignIn />} />
        <Route path="/signin" element={<SignIn />} />
        <Route path="/signup" element={<SignUp />} />
        <Route path="/home" element={[<NavBar />, <Home />]} />
        <Route path="/:uuid" element={<Url />} />
    </Routes>
    </BrowserRouter>
  );
}

export default App;

import './App.css'
import {getList, get, create, update, remove} from "./clients.api.js";

function App() {

  const consoleTable = (res) => {
    console.table(res)
  }

  const getClientList = async ()  => {
    await getList().then(res => consoleTable(res.data))
  }

  const getClient = async () => {
    await get(1).then(res => consoleTable(res.data))
  }

  const createClient = async () => {
    await create({
      name: 'John Doe',
      TIN_or_KPP: '1234567890',
      address: '123 Main Street',
      BIK: '012345678',
      checking_account: '9876543210',
      correspondent_account: '98765432109876543210'
    }).then(res => consoleTable(res.data))
  }

  const updateClient = async () => {
    await update(1, {
      name: 'NEW John Doe',
      TIN_or_KPP: '123NEW7890',
      address: 'NEW 123 Main Street',
      BIK: '012NEW678',
      checking_account: '987NEW3210',
      correspondent_account: '987NEW32109876543210'
    }).then(res => console.log(res.status))
  }

  const deleteClient = async () => {
    await remove(6).then(res => console.log(res.status))
  }

  return (
    <div>
      <button onClick={getClientList}>GET CLIENT LIST</button>
      <button onClick={getClient}>GET CLIENT</button>
      <button onClick={createClient}>CREATE CLIENT</button>
      <button onClick={updateClient}>UPDATE CLIENT</button>
      <button onClick={deleteClient}>DELETE CLIENT</button>
    </div>
  )
}

export default App

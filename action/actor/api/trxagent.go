// Copyright 2017~2022 The Bottos Authors
// This file is part of the Bottos Chain library.
// Created by Rocket Core Team of Bottos.

//This program is free software: you can distribute it and/or modify
//it under the terms of the GNU General Public License as published by
//the Free Software Foundation, either version 3 of the License, or
//(at your option) any later version.

//This program is distributed in the hope that it will be useful,
//but WITHOUT ANY WARRANTY; without even the implied warranty of
//MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//GNU General Public License for more details.

//You should have received a copy of the GNU General Public License
// along with bottos.  If not, see <http://www.gnu.org/licenses/>.

/*
 * file description:  trx agent
 * @Author:
 * @Date:   2017-12-06
 * @Last Modified by:
 * @Last Modified time:
 */

package apiactor

import (
	"fmt"
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/bottos-project/core/common/types"
)

var trxactorPid *actor.PID

func SetTrxActorPid(tpid *actor.PID) {
	trxactorPid = tpid
}

func PushTransaction(trx uint64) error {

	fmt.Println("exec push transaction: ", trx)

	pushTrxReq := &types.Transaction{
		RefBlockNum: 11,
		Sender:      22,
		Action:      trx,
	}
	/*
		trxactorPid.Tell(pushTrxReq)

		f := trxactorPid.RequestFuture(pushTrxReq, 5000*time.Millisecond)
		es, err := f.Result() // waits for pid to reply

		fmt.Println("this is es err", es, err)
	*/

	result, err := trxactorPid.RequestFuture(pushTrxReq, 500*time.Millisecond).Result() // await result

	fmt.Println("rusult is =======", result, err)

	//trxactorPid.Tell(pushTrxReq)

	fmt.Println("exec push transaction done !!!!!!", trx)

	return nil
}